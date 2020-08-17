import { Card } from '@material-ui/core';
import IconButton from '@material-ui/core/IconButton';
import AddIcon from '@material-ui/icons/Add';
import React from "react";
import makeStyles from "@material-ui/core/styles/makeStyles";
import CardContent from "@material-ui/core/CardContent";
import { connect } from 'react-redux';

import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import withStyles from "@material-ui/core/styles/withStyles";

import {Creators as nightsActions} from '../../store/reducers/nights';

const styles = (theme) => {
    return {
        root: {
            maxWidth: 345,
        },
        addCard: {
            width: '300px',
            justifyContent: "center"
        },
        alignItemsAndJustifyContent: {
            width: 300,
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center'
        }
    }
}


export class AddNightCard extends React.Component {
    constructor(props) {
        super(props);
        this.state = {open: false, date: '', location: ''};
        this.handleDateChange = this.handleDateChange.bind(this);
        this.handleLocationChange = this.handleLocationChange.bind(this);
        this.handleClickOpen = this.handleClickOpen.bind(this);
        this.handleClose = this.handleClose.bind(this);
        this.onClickAdd = this.onClickAdd.bind(this);
    }
    handleDateChange(event) {    this.setState({...this.state, date: event.target.value});  }
    handleLocationChange(event) {    this.setState({...this.state, location: event.target.value});  }
    handleClickOpen (event){
        this.setState({...this.state, open:true});
    }

    handleClose (){
        this.setState({...this.state, open:false});
    }
    onClickAdd(){
        this.props.addNight({date: this.state.date, location: this.state.location, group: this.props.group.id, token: this.props.user.token});
        this.setState({...this.state, open:false});
    }

    render() {
        const { classes } = this.props;
        return (
            <Card className={classes.alignItemsAndJustifyContent}>
                <CardContent>
                    <div>
                        <Button variant="outlined" color="primary" onClick={this.handleClickOpen}>
                            <AddIcon fontSize={"large"} />
                        </Button>
                        <Dialog open={this.state.open} onClose={this.handleClose} aria-labelledby="form-dialog-title">
                            <DialogTitle id="form-dialog-title">Add Movie</DialogTitle>
                            <DialogContent>
                                <DialogContentText>
                                    Enter the movie night information
                                </DialogContentText>
                                <TextField
                                    autoFocus
                                    margin="dense"
                                    id="location"
                                    label="Location"
                                    type="text"
                                    fullWidth
                                    onChange={this.handleLocationChange}
                                />
                                <TextField
                                    autoFocus
                                    margin="dense"
                                    id="date"
                                    label="Date"
                                    type="date"
                                    fullWidth
                                    onChange={this.handleDateChange}
                                />
                            </DialogContent>
                            <DialogActions>
                                <Button onClick={this.handleClose} color="primary">
                                    Cancel
                                </Button>
                                <Button onClick={this.onClickAdd} color="primary">
                                    Add
                                </Button>
                            </DialogActions>
                        </Dialog>
                    </div>
                </CardContent>

            </Card>
        );
    }


}

const mapDispatchToProps = (dispatch)=>({
    addNight: (movie)=> dispatch(nightsActions.add_night_request(movie))
});

export default connect(undefined, mapDispatchToProps)(withStyles(styles)(AddNightCard));