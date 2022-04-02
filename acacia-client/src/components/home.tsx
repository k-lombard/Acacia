import React, { useState, useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'
import { AnyAction } from 'redux'
import { RootState } from '../store/store'
import { PURGE } from 'redux-persist'
import Button from '@mui/material/Button'
import { setLoading } from '../store/Acacia-actions'
import LayeredMap from './map';

export default function Acacia(this: any) {
  const forceUpdate = useForceUpdate()
  const dispatch = useDispatch<ThunkDispatch<{}, {}, AnyAction>>()
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



  return (
    <div>
      <h1>Acacia</h1>
      <Button variant="contained" color="primary" onClick={purgeStore}>New Images</Button>
      <LayeredMap />
    </div>
  )
}  
