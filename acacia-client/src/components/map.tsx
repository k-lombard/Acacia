/// app.js
import React from 'react';
import DeckGL from '@deck.gl/react';
import {LineLayer} from '@deck.gl/layers';
import {StaticMap} from 'react-map-gl';

// Set your mapbox access token here
const MAPBOX_ACCESS_TOKEN = 'pk.eyJ1IjoibWlydmluZSIsImEiOiJjbDFpNXhva2MxbHAxM2pxaXEwdjEzMmw3In0.cmMa2dAsNf_4TiUjpD9HiQ';


// Viewport settings
const INITIAL_VIEW_STATE = {
  longitude: 33.772163578,
  latitude: -84.390165106,
  zoom: 13,
  pitch: 0,
  bearing: 0
};

// Data to be used by the LineLayer
const data = [
  {sourcePosition: [-122.41669, 37.7853], targetPosition: [-122.41669, 37.781]}
];

export default function Map() {
  return (
    <DeckGL
      initialViewState={INITIAL_VIEW_STATE}
      controller={true} >
      <LineLayer id="line-layer" data={data} />
    </DeckGL>
  );
}