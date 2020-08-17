import React from 'react';
import Header from './Header';
import Button from '@material-ui/core/Button';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import Image from 'material-ui-image';
import theater from '../assets/img/thater.jpg';
import {Link} from 'react-router-dom';


const LandingPage = () => (
  <div className="landingPagebodyComponent">

<br/>
<Typography variant="display3" gutterBottom align="center">
        Welcome to Based Movie Nights
      </Typography>
    
   <Grid container spacing={24} >
        <Grid item xs={6}>
          <Typography gutterBottom align="left" style={{paddingLeft:20}}>
        {`
         Can't decide which movie to watch? Say no more! Based Movie Nights Will take care of this problem!
        `}
        <Link to="/login"> <Button color="primary"  align="left" style={{marginLeft:20}}>
        Get Started
      </Button></Link>
      </Typography>
        </Grid>
        <Grid item xs={12}>
         
        </Grid>
       {/* <Grid item xs={6}>
          <Paper >xs=6</Paper>
        </Grid>*/}
        </Grid>

     <Grid container spacing={24} >
        <Grid item xs={12} md={12}>
            <Grid item xs={6} style={{display: 'flex',
                alignItems: 'center',
                justifyContent: 'center'}}>
                <Image
                    src={theater}
                    color="inherit" style={{height:40}} imageStyle={{ width: '50', height: '75%' }} />
            </Grid>

          </Grid>
          </Grid>

  </div>
);

export default LandingPage;