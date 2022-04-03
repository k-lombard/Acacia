import React from 'react';
import logo from './logo.svg';
import './App.css';
import { Provider } from 'react-redux';
import Acacia from './components/home';
import LayeredMap from './components/map';
import { PersistGate } from 'redux-persist/integration/react';
import { persistStore } from 'redux-persist';
import store from './store/store';
import { ChakraProvider } from '@chakra-ui/react'
const persistor = persistStore(store);


function App() {
  return (
    <div className="App">
      <Provider store={store}>
          <Acacia/>
      </Provider>
    </div>
  );
}

export default App;
