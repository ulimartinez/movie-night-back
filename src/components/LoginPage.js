import React from 'react';
import { connect } from 'react-redux';
import {Types as authTypes, Creators as authActions} from '../store/reducers/auth';
import Typography from '@material-ui/core/Typography';
import Paper from '@material-ui/core/Paper';
import Input from '@material-ui/core/Input';
import InputLabel from '@material-ui/core/InputLabel';
import InputAdornment from '@material-ui/core/InputAdornment';
import FormControl from '@material-ui/core/FormControl';
import {Link, Redirect} from 'react-router-dom';
import TextField from '@material-ui/core/TextField';
import Grid from '@material-ui/core/Grid';
import AccountCircle from '@material-ui/icons/AccountCircle';
import Key from '@material-ui/icons/VpnKey';
import Button from '@material-ui/core/Button';
import withStyles from "@material-ui/core/styles/withStyles";

const styles = (theme) => {
    return {
        button: {
            color: '#0766b9'
        },
    }
}

export class LoginPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {username: '', password: ''};
        this.handleUsernameChange = this.handleUsernameChange.bind(this);
        this.handlePasswordChange = this.handlePasswordChange.bind(this);
    }

    handleUsernameChange(event) {
        this.setState({...this.state, username: event.target.value});
    }

    handlePasswordChange(event) {
        this.setState({...this.state, password: event.target.value});
    }

    OnClickLogin = () => {
        this.props.startLogin({email: this.state.username, password: this.state.password});

    };

    render() {
        const {
            classes,
            connected,
            error,
            errorMsg,
            loading,
            onChangePwdField,
            onSubmit
        } = this.props;

        if (connected) {
            return <Redirect to="/home"/>;
        }
        return (
            <div className="login-page-class">


                <Paper className="loginPaper">
                    <div className="loginheaderpart">
                        <Typography variant="display3" gutterBottom className="loginpageheader">
                            Login
                        </Typography>
                    </div>
                    <Typography variant="headline" component="h3">
                        Login to your account
                    </Typography>
                    <form>
                        <div className="loginformgroup">

                            <AccountCircle/>

                            <TextField id="input-username" label="Username" value={this.state.username}
                                       onChange={this.handleUsernameChange}/>

                        </div>
                        <div className="loginformgroup">


                            <Key/>

                            <TextField type="password" id="input-password" label="Password" value={this.state.password}
                                       onChange={this.handlePasswordChange}/>

                        </div>
                    </form>

                    <Button variant="raised" color="blue" onClick={this.OnClickLogin}><Typography variant="button"
                                                                                                     gutterBottom
                                                                                                     className={classes.button}>
                        Login
                    </Typography></Button>
                    <Link to={"/register"} variant="raised" color="blue"><Typography variant="button" gutterBottom className={classes.button}>
                        Register
                    </Typography></Link>
                </Paper>
            </div>
        );
    }
}
const mapStateToProps = state => ({
    connected: state.auth.connected,
    loading: state.auth.fetching
});



const mapDispatchToProps = (dispatch)=>({
  startLogin: (user)=> dispatch(authActions.auth_request(user))
});


export default connect(mapStateToProps, mapDispatchToProps)(withStyles(styles)(LoginPage));
