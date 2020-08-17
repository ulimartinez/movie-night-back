import React from 'react';
import {render} from 'react-dom';
import { MuiThemeProvider, createMuiTheme } from '@material-ui/core/styles';
import { Provider } from 'react-redux';
import HeaderAppBar from './components/Header';
import AppRouter from './routers/AppRouter';
import './styles/styles.css';
import {theme} from './theme/theme';
import configureStore from './store/config/configureStore';


const store = configureStore({});



const App = () => (
  <Provider store={store}>
  <MuiThemeProvider theme={theme}>
    <AppRouter />
  </MuiThemeProvider>
  </Provider>
);


render(<App />, document.getElementById('app'));