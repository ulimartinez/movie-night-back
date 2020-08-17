
import React from 'react';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import PropTypes from "prop-types";
import HomeIcon from '@material-ui/icons/Home';
import MovieIcon from '@material-ui/icons/Movie';
import Brightness3Icon from '@material-ui/icons/Brightness3';
import GroupIcon from '@material-ui/icons/Group';

/* import your desired icon from material-ui icons library */
import {NavLink} from 'react-router-dom';


export const publicNavs = [
    {
        url:'/home',
        name:'Home',
        icon:<HomeIcon/>
    },
    {
        url:'/movies',
        name:'Movies',
        icon:<MovieIcon/>
    },
    {
        url:'/nights',
        name:'Nights',
        icon:<Brightness3Icon/>
    },
    {
        url:'/groups',
        name:'Groups',
        icon:<GroupIcon/>
    }
];



export default  ()=>(
publicNavs.map((navItem)=>{
return <NavLink to={navItem.url}  className="NavLinkItem" key={navItem.url} activeClassName="NavLinkItem-selected"> <List component="nav" >  <ListItem button>
          <ListItemIcon className="innernavitem"> 
 {navItem.icon}
          </ListItemIcon>
          <ListItemText primary={navItem.name} className="innernavitem" color="black"/>
        </ListItem></List> </NavLink>
})


     

);





