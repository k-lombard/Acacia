import React, { useState, useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'
import { AnyAction } from 'redux'
import { RootState } from '../store/store'
import { PURGE } from 'redux-persist'
import Button from '@mui/material/Button'
import { setLoading } from '../store/Acacia-actions'
import { AppBar, Container, Toolbar, Typography, Box, IconButton, Menu, MenuItem, Tooltip, CssBaseline } from '@mui/material'
import { MenuIcon, Link } from '@chakra-ui/react'
import LayeredMap from './map'
import Images from './images'
import Avatar from '@mui/material/Avatar';
import { getAllSentries } from '../store/Acacia-actions'
import { getImagesBySentryId } from '../store/Acacia-actions'
import { BrowserRouter, Routes, Route } from "react-router-dom";

export default function Acacia(this: any) {
  const forceUpdate = useForceUpdate()
  const dispatch = useDispatch<ThunkDispatch<{}, {}, AnyAction>>()
  const [anchorElNav, setAnchorElNav] = React.useState<null | HTMLElement>(null);
  const [anchorElUser, setAnchorElUser] = React.useState<null | HTMLElement>(null);
  const [currentImages, setCurrentImages] = useState([]);
  const [currentSentries, setCurrentSentries] = useState([]);
  const pages = ['Home', 'Map', 'Images', 'Sentries', 'About'];
  const settings = ['Profile', 'Account', 'Dashboard', 'Logout'];
  function useForceUpdate(){
    const [value, setValue] = useState(0); // integer state
    return () => setValue(value => value + 1); // update the state to force render
  }


  const purgeStore = (evt: React.MouseEvent) => { 
    evt.preventDefault();
    dispatch({ 
        type: PURGE,
        key: "root",    // Whatever you chose for the "key" value when initialising redux-persist in the **persistCombineReducers** method - e.g. "root"
        result: () => null              // Func expected on the submitted action. 
    })  
    dispatch(setLoading(true))  
  }

  useEffect(() => {
    if (currentImages.length === 0 || currentSentries.length === 0) {
      dispatch(getAllSentries()).then(res => {
        setCurrentSentries(res.sentries);
        for (let sentry of res.sentries) {
          dispatch(getImagesBySentryId(sentry.id)).then(res2 => {
            setCurrentImages(res2.images);
            console.log(res2.images)
          })
        }
      })
    }
  }, [])

  const handleOpenNavMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorElNav(event.currentTarget);
  };
  const handleOpenUserMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorElUser(event.currentTarget);
  };

  const handleCloseNavMenu = () => {
    setAnchorElNav(null);
  };

  const handleCloseUserMenu = () => {
    setAnchorElUser(null);
  };

  return (
    <div>
     <AppBar position="static" style={{backgroundColor:"gray", height: "10%", maxHeight: "10%"}}>
      <CssBaseline />
      <Toolbar>
      <img style={{height: "64px", width: "84px"}}src={require("./acacia.jpeg")}></img>
          <div style={{display: "flex", justifyContent: "space-between", flexDirection: "row"}}>
            <Link style={{margin: "10px", textDecoration: "none"}} href="/map" color="white">
              Map
            </Link>
            <Link style={{margin: "10px", textDecoration: "none"}} href="/images" color="white">
              Images
            </Link>
          </div>
      </Toolbar>
    </AppBar>

    <div style={{ height: '80vh', width: '100vw' }}>
    <Routes>
      <Route
          path={"/images"}
          element={
            <Images images={currentImages}
            />}
        />
        <Route
          path={"/map"}
          element={
            <LayeredMap
            />
          }
        />
        <Route
          path={"/"}
          element={
            <LayeredMap
            />
          }
        />
      </Routes>
    </div>
    </div>
  )
}  
