import React from 'react';
import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import AccountCircle from '@material-ui/icons/AccountCircle';
import Key from '@material-ui/icons/VpnKey';
import Button from '@material-ui/core/Button';
import Snackbar from '@material-ui/core/Snackbar';
import IconButton from '@material-ui/core/IconButton';
import CloseIcon from '@material-ui/icons/Close';
import Typography from '@material-ui/core/Typography';
import fusioncharts from 'fusioncharts';
// Load the charts module
import charts from 'fusioncharts/fusioncharts.charts';
import ReactFC from 'react-fusioncharts';
import {connect} from "react-redux";

import {Creators as moviesActions} from "../store/reducers/movies";

import MovieCard from '../components/Cards/MovieCard';
import AddCard from "./Cards/AddMovieCard";

class MoviesPage extends React.Component {
    constructor(props) {
        super(props);
        this.handleSnackClose = this.handleSnackClose.bind(this);
    }
    componentDidMount() {
        this.props.fetchMovies({id: this.props.group.id, user: this.props.user});
    }
    handleSnackClose(event, reason){
        if(reason === 'clickaway'){
            return;
        }
        this.props.dismissSnack();
    }
    render() {
        const movies_cards = [];
        if(this.props.movies){
            for(let i = 0; i < this.props.movies.length; i++){
                movies_cards.push(<MovieCard movie={{movie:this.props.movies[i]}} user={this.props.user} key={i} />);
            }
        }
        return (
            <React.Fragment>

                <Grid container spacing={24}>
                    <Grid item xs={12} style={{marginTop:10}}>
                        <Paper className="homepagepaper">
                            <Typography variant="body1" gutterBottom align="center">
                                This is a list of your movies.
                            </Typography>

                        </Paper>
                    </Grid>

                </Grid>
                <Grid container spacing={24}>
                    {movies_cards}
                    <AddCard group={this.props.group} user={this.props.user} getMovies={this.props.fetchMovies} />
                </Grid>
                <Snackbar
                    anchorOrigin={{
                        vertical: 'bottom',
                        horizontal: 'left',
                    }}
                    open={this.props.error || this.props.snack}
                    autoHideDuration={6000}
                    onClose={this.handleSnackClose}
                    message={this.props.snack || this.props.errorMessage.date || ''}
                    action={
                        <React.Fragment>
                            <IconButton size="small" aria-label="close" color="inherit" onClick={this.handleSnackClose}>
                                <CloseIcon fontSize="small" />
                            </IconButton>
                        </React.Fragment>
                    }
                />
            </React.Fragment>
        );
    }
}
const mapStateToProps = state => ({
    connected: state.auth.connected,
    loading: state.auth.fetching,
    user: state.auth.user,
    group: state.groups.currentGroup,
    movies: state.movies.submissions,
    snack: state.movies.snack,
    error: state.movies.error,
    errorMessage: state.movies.errors? state.movies.errors : ''
});

const mapDispatchToProps2 = (dispatch)=>({
    fetchMovies: (data)=> dispatch(moviesActions.get_movies_request(data)),
    dismissSnack: () => dispatch(moviesActions.dismiss_snack())
});


export default connect(mapStateToProps, mapDispatchToProps2)(MoviesPage);