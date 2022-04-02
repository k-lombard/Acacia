/// app.js
import React from 'react';
import DeckGL from 'deck.gl';
import { Map } from "react-map-gl";

export default function LayeredMap() {
  return (
    <Map
        initialViewState={{
          longitude: -100,
          latitude: 40,
          zoom: 3.5
        }}
        style={{width: '100vw', height: '100vh'}}
        mapStyle="mapbox://styles/mapbox/streets-v9"
        mapboxAccessToken="pk.eyJ1IjoibWlydmluZSIsImEiOiJjbDFpNXhva2MxbHAxM2pxaXEwdjEzMmw3In0.cmMa2dAsNf_4TiUjpD9HiQ"
      />
  )
}