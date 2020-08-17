import React from 'react';
import { BrowserRouter, Route, Switch, Link, NavLink } from 'react-router-dom';
import LandingPage from '../components/LandingPage';
import AboutPage from '../components/AboutPage';
import HomePage from '../components/HomePage';
import ContactPage from '../components/ContactPage';
import FAQPage from '../components/FAQPage';
import NotFoundPage from '../components/NotFoundPage';
import ServicesPage from '../components/ServicesPage';

import PublicRoute from './PublicRouter';
import PrivateRoute from './PrivateRouter';
import LoginPage from '../components/LoginPage';
import GetStartedPage from '../components/GetStartedPage';
import ProductsPage from '../components/ProductsPage';
import FormsPage from '../components/FormsPage';
import MoviesPage from "../components/MoviesPage";
import NightsPage from "../components/NightsPage";
import RegisterPage from "../components/RegisterPage";
import GroupsPage from "../components/GroupsPage";





const AppRouter = () => (
  <BrowserRouter>
    <div>
      
      <Switch>
        
        <PublicRoute path="/" component={LandingPage} exact={true} />
         <PrivateRoute path="/home" component={HomePage} />
          <PrivateRoute path="/movies" component={MoviesPage} />
          <PrivateRoute path="/nights" component={NightsPage} />
          <PrivateRoute path="/groups" component={GroupsPage} />
       
        <Route path="/login" component={LoginPage}  />

        
        <PublicRoute path="/register" component={RegisterPage} />

        <Route component={NotFoundPage} />
      </Switch>
      
    </div>
  </BrowserRouter>
);

export default AppRouter;