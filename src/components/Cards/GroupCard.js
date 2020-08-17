import { Card } from '@material-ui/core';
import React from "react";
import CardActions from "@material-ui/core/CardActions";
import Button from "@material-ui/core/Button";
import CardContent from "@material-ui/core/CardContent";
import Typography from "@material-ui/core/Typography";
import CardMedia from "@material-ui/core/CardMedia";
import {connect} from "react-redux";
import withStyles from "@material-ui/core/styles/withStyles";

import {Creators as groupsActions} from "../../store/reducers/groups";

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
    },
    greenBG: {
        'background-color': 'rgb(25, 118, 210)'
    }
});

export class GroupCard extends React.Component{
    constructor(props) {
        super(props);
        this.handleAssign = this.handleAssign.bind(this);
    }
    handleAssign = () => {
        this.props.changeGroup({groupId: this.props.group.id, token: this.props.user.token})
    }
    render() {
        const {classes} = this.props;
        const current = (<Typography color="textSecondary" gutterBottom className={classes.greenBG}>
            Current Group:
        </Typography>);
        const normal = (<Typography color="textSecondary" gutterBottom>
            Group:
        </Typography>);
        /*const movie = (<CardMedia
            className={classes.media}
            image={this.props.movie.movie.Poster}
            title={this.props.movie.movie.title}
        />);
        */
        return (
            <Card className={classes.movieCard}>
                <CardContent>
                    {this.props.currentGroup.id === this.props.group.id? current: normal}
                    <Typography variant="h5" component="h2">
                        Name: {this.props.group.name}
                    </Typography>
                    <Typography variant="h5" component="h2">
                        ID: {this.props.group.id}
                    </Typography>
                </CardContent>
                <CardActions>
                    <Button
                        size="small"
                        onClick={this.handleAssign}
                    >
                        Select
                    </Button>
                </CardActions>
            </Card>
        );
    }
}

const mapDispatchToProps = (dispatch)=>({
    changeGroup: (data)=> dispatch(groupsActions.set_groups_request(data)),
})
export default connect(undefined, mapDispatchToProps)(withStyles(styles)(GroupCard));