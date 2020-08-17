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

import {Creators as nightsActions} from "../store/reducers/nights";

import NightCard from '../components/Cards/NightCard';
import AddNightCard from "../components/Cards/AddNightCard";

class NightsPage extends React.Component {
    componentDidMount() {
        this.props.fetchNights({id: this.props.group.id, user: this.props.user});
    }

    render() {
        const nights_cards = [];
        if(this.props.nights){
            for(let i = 0; i < this.props.nights.length; i++){
                nights_cards.push(<NightCard group={this.props.group} night={this.props.nights[i]} user={this.props.user} key={i} />);
            }
        }
        return (
            <React.Fragment>

                <Grid container spacing={24}>
                    <Grid item xs={12} style={{marginTop:10}}>
                        <Paper className="homepagepaper">
                            <Typography variant="body1" gutterBottom align="center">
                                This is a list of your movie nights.
                            </Typography>

                        </Paper>
                    </Grid>

                </Grid>
                <Grid container spacing={24}>
                    {nights_cards}
                    <AddNightCard group={this.props.group} user={this.props.user} getMovies={this.props.fetchNights} />
                </Grid>
            </React.Fragment>
        );
    }
}
const mapStateToProps = state => ({
    loading: state.auth.fetching,
    user: state.auth.user,
    group: state.groups.currentGroup,
    nights: state.nights.nights
});

const mapDispatchToProps = (dispatch)=>({
    fetchNights: (data)=> dispatch(nightsActions.get_nights_request(data))
});


export default connect(mapStateToProps, mapDispatchToProps)(NightsPage);