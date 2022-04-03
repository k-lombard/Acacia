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

const data = [{name: 'Klaus College of Computing', address: '266 Ferst Dr NW, Atlanta, GA 30332', exits: 4214, coordinates: [ -84.39579759751776,33.77715462850155]}];

function App() {
  return (
    <div className="App">
      test
      <Provider store={store}>
        <PersistGate loading={null} persistor={persistor}>
          <Acacia/>
          <LayeredMap data={data}/>
        </PersistGate>
      </Provider>
    </div>
  );
}

export default App;
