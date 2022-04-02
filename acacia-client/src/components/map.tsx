/// app.js
import React from 'react';
import DeckGL from 'deck.gl';
import {LineLayer} from '@deck.gl/layers';
import {IconLayer} from '@deck.gl/layers';
import { Map } from "react-map-gl";

// Set your mapbox access token here
const MAPBOX_ACCESS_TOKEN = 'pk.eyJ1IjoibWlydmluZSIsImEiOiJjbDFpNXhva2MxbHAxM2pxaXEwdjEzMmw3In0.cmMa2dAsNf_4TiUjpD9HiQ';

// Viewport settings
const INITIAL_VIEW_STATE = {
  longitude: -84.39634476811625,
  latitude: 33.77565864190424,
  zoom: 15,
  pitch: 0,
  bearing: 0
};

const data = [
  {sourcePosition: [-84.390165106, 33.772163578], targetPosition: [-122.41669, 37.781]}
];

const ICON_MAPPING = {
  marker: {x: 0, y: 0, width: 128, height: 128, mask: true}
};

export default function LayeredMap({data}) {
  const layer = new IconLayer({
    id: 'icon-layer',
    data,
    pickable: true,
    // iconAtlas and iconMapping are required
    // getIcon: return a string
    iconAtlas: 'https://raw.githubusercontent.com/visgl/deck.gl-data/master/website/icon-atlas.png',
    iconMapping: ICON_MAPPING,
    getIcon: d => 'marker',
  
    sizeScale: 15,
    getPosition: d => d.coordinates,
    getSize: d => 5,
    getColor: d => [Math.sqrt(d.exits), 140, 0]
  });

  return (
    <DeckGL
      initialViewState={INITIAL_VIEW_STATE}
      controller={true}
      layers={layer}  
    >
      <Map
          style={{width: '100vw', height: '100vh'}}
          mapStyle="mapbox://styles/mapbox/streets-v9"
          mapboxAccessToken={MAPBOX_ACCESS_TOKEN}
        />
    </DeckGL>
  )
}