/// app.js
import React from 'react';
import DeckGL from 'deck.gl';
import {LineLayer} from '@deck.gl/layers';
import {IconLayer} from '@deck.gl/layers';
import { Map } from "react-map-gl";

// Set your mapbox access token here
const MAPBOX_ACCESS_TOKEN = 'pk.eyJ1IjoibWlydmluZSIsImEiOiJjbDFpNXhva2MxbHAxM2pxaXEwdjEzMmw3In0.cmMa2dAsNf_4TiUjpD9HiQ';
const data2 = [{name: 'Klaus College of Computing', alias: "SentryOne", address: '266 Ferst Dr NW, Atlanta, GA 30332', exits: 4214, sentry_id: "e31e7bf3-64b1-4aa6-a02e-aac0b6efe1e0", coordinates: [ -84.39579759751776,33.77715462850155]}];

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

export default function LayeredMap() {
  const layer = new IconLayer({
    id: 'icon-layer',
    data: data2,
    pickable: true,
    // iconAtlas and iconMapping are required
    // getIcon: return a string
    iconAtlas: 'https://raw.githubusercontent.com/visgl/deck.gl-data/master/website/icon-atlas.png',
    iconMapping: ICON_MAPPING,
    getIcon: d => 'marker',
  
    sizeScale: 8,
    getPosition: d => d.coordinates,
    getSize: d => 5,
    getColor: d => [Math.sqrt(d.exits), 380, 80]
  });

  return (
    <DeckGL
      initialViewState={INITIAL_VIEW_STATE}
      controller={true}
      layers={layer}
      getTooltip={({object}) => object && `Alias: ${object.alias}\n SentryId: ${object.sentry_id}\n${object.name}\n${object.address}`} 
      width={'100%'} 
      height={'100%'}
      style={{marginTop: "64px"}}
    >
      <Map
          mapStyle="https://basemaps.cartocdn.com/gl/dark-matter-gl-style/style.json"
          mapboxAccessToken={MAPBOX_ACCESS_TOKEN}
        />
    </DeckGL>
  )
}