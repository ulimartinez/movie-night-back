import { Card } from '@material-ui/core';
import React from "react";
import CardActions from "@material-ui/core/CardActions";
import Button from "@material-ui/core/Button";
import CardContent from "@material-ui/core/CardContent";
import Typography from "@material-ui/core/Typography";
import CardMedia from "@material-ui/core/CardMedia";
import {connect} from "react-redux";
import withStyles from "@material-ui/core/styles/withStyles";

import {Creators as NightsActions} from "../../store/reducers/nights";

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

export class NightCard extends React.Component{
    constructor(props) {
        super(props);
        this.handleAssign = this.handleAssign.bind(this);
    }
    handleAssign = () => {
       this.props.selectMovie({groupId: this.props.group.id, nightId: this.props.night.id, token: this.props.user.token})
    }
    render() {
        const {classes} = this.props;
        /*const movie = (<CardMedia
            className={classes.media}
            image={this.props.movie.movie.Poster}
            title={this.props.movie.movie.title}
        />);
        */
        return (
            <Card className={classes.movieCard}>
                <CardContent>
                    <Typography color="textSecondary" gutterBottom>
                        Movie Night:
                    </Typography>
                    <Typography variant="h5" component="h2">
                        Date: {this.props.night.date}
                    </Typography>
                    <Typography variant="h5" component="h2">
                        Host: {this.props.night.Host.username}
                    </Typography>
                    <Typography variant="h5" component="h2">
                        Location: {this.props.night.location}
                    </Typography>
                    <Typography variant="h5" component="h2">
                        Movie: {this.props.night.Movie.title}
                    </Typography>
                </CardContent>
                <CardActions>
                    <Button
                        size="small"
                        onClick={this.handleAssign}
                    >
                        Select Movie
                    </Button>
                </CardActions>
            </Card>
        );
    }
}


const mapDispatchToProps = (dispatch)=>({
    selectMovie: (data)=> dispatch(NightsActions.assign_movie_request(data)),
})

export default connect(undefined, mapDispatchToProps)(withStyles(styles)(NightCard));