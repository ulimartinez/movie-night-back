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
import IconButton from "@material-ui/core/IconButton";
import InputBase from "@material-ui/core/InputBase";
import Paper from "@material-ui/core/Paper";
import OutlinedInput from "@material-ui/core/OutlinedInput";

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
    input: {
        marginLeft: theme.spacing(1),
        flex: 1,
    },
});

export class JoinGroupCard extends React.Component{
    constructor(props) {
        super(props);
        this.handleJoin = this.handleJoin.bind(this);
        this.handleJoinIDChange = this.handleJoinIDChange.bind(this);
        this.state = {...this.state, joinId: ''};
    }
    handleJoin = () => {
        this.props.joinGroup({groupId: this.state.joinId, token: this.props.user.token})
    }
    handleJoinIDChange(event) {    this.setState({...this.state, joinId: event.target.value});  }
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
                    <Paper component="form" className={classes.root}>
                        <OutlinedInput
                            className={classes.input}
                            placeholder="Group ID"
                            onChange={this.handleJoinIDChange}
                        />
                        <Button
                        onClick={this.handleJoin}>
                            Join
                        </Button>
                    </Paper>
                </CardContent>
            </Card>
        );
    }
}

const mapDispatchToProps = (dispatch)=>({
    joinGroup: (data)=> dispatch(groupsActions.join_group_request(data)),
})
export default connect(undefined, mapDispatchToProps)(withStyles(styles)(JoinGroupCard));