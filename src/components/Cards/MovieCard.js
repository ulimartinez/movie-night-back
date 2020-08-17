import { Card } from '@material-ui/core';
import React from "react";
import CardActions from "@material-ui/core/CardActions";
import Button from "@material-ui/core/Button";
import CardContent from "@material-ui/core/CardContent";
import Typography from "@material-ui/core/Typography";
import CardMedia from "@material-ui/core/CardMedia";
import {connect} from "react-redux";
import withStyles from "@material-ui/core/styles/withStyles";

import {Creators as moviesActions} from "../../store/reducers/movies";

const styles = (theme) => ({
    root: {
        maxWidth: 345,
    },
    media: {
        height: 0,
        paddingTop: '100%',
        'background-size': 'contain'
    },
    movieCard: {
        width: '300px'
    }
});

export class MovieCard extends React.Component{
    constructor(props) {
        super(props);
        this.handleVote = this.handleVote.bind(this);
    }
    handleVote = () => {
       this.props.voteMovie({movieId: this.props.movie.movie.id, token: this.props.user.token})
    }
    render() {
        const {classes} = this.props;
        return (
            <Card className={classes.movieCard}>
                <CardContent>
                    <Typography color="textSecondary" gutterBottom>
                        Movie:
                    </Typography>
                    <Typography variant="h5" component="h2">
                        {this.props.movie.movie.Title}
                    </Typography>
                    <CardMedia
                        className={classes.media}
                        image={this.props.movie.movie.Poster}
                        title={this.props.movie.movie.title}
                    />
                </CardContent>
                <CardActions>
                    <Button
                        size="small"
                        onClick={this.handleVote}
                    >
                        Vote
                    </Button>
                </CardActions>
            </Card>
        );
    }
}

const mapDispatchToProps = (dispatch)=>({
    voteMovie: (data)=> dispatch(moviesActions.vote_request(data)),
})

export default connect(undefined, mapDispatchToProps)(withStyles(styles)(MovieCard));