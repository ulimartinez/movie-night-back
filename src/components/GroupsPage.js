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

import {Creators as groupsActions} from "../store/reducers/groups";

import AddGroupCard from "../components/Cards/AddGroupCard";
import GroupCard from "./Cards/GroupCard";
import JoinGroupCard from "./Cards/JoinGroupCard";

class GroupsPage extends React.Component {
    componentDidMount() {
        this.props.fetchGroups(this.props.user);
    }

    render() {
        const groups_cards = [];
        if(this.props.groups){
            for(let i = 0; i < this.props.groups.length; i++){
                groups_cards.push(<GroupCard group={this.props.groups[i]} user={this.props.user} currentGroup={this.props.currentGroup} key={i} />);
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
                    {groups_cards}
                    <JoinGroupCard user={this.props.user} />
                    <AddGroupCard group={this.props.group} user={this.props.user} getMovies={this.props.fetchNights} />
                </Grid>
            </React.Fragment>
        );
    }
}
const mapStateToProps = state => ({
    loading: state.auth.fetching,
    user: state.auth.user,
    groups: state.groups.groups,
    nights: state.nights.nights,
    currentGroup: state.groups.currentGroup
});

const mapDispatchToProps = (dispatch)=>({
    fetchGroups: (user)=> dispatch(groupsActions.groups_request(user))
});


export default connect(mapStateToProps, mapDispatchToProps)(GroupsPage);