/// app.js
import React from 'react';
import DeckGL from 'deck.gl';
import {LineLayer} from '@deck.gl/layers';
import {IconLayer} from '@deck.gl/layers';
import { Map } from "react-map-gl";
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import Divider from '@mui/material/Divider';
import ListItemText from '@mui/material/ListItemText';
import ListItemAvatar from '@mui/material/ListItemAvatar';
import ListItemButton from '@mui/material/ListItemButton';
import Avatar from '@mui/material/Avatar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';



export default function Images({images}) {
    const [open, setOpen] = React.useState(false);
    const [currImg, setCurrImg] = React.useState("");
    const handleClickOpen = (img: string) => {
        setCurrImg(img);
        setOpen(true);
    };

    const handleClose = () => {
        setOpen(false);
    };


  return (
      <React.Fragment>
    <Dialog
    open={open}
    onClose={handleClose}
    aria-labelledby="alert-dialog-title"
    aria-describedby="alert-dialog-description"
  >
        <DialogContent>
          <img src={currImg}></img>
        </DialogContent>
      <DialogActions>
          <Button onClick={handleClose}>Close</Button>
        </DialogActions>
      </Dialog>
    <List dense style={{ width: '100%', height: "100%", justifyContent: "center", justifySelf: "center", alignItems: "center"}}>
      {images.map((img) => {
        const labelId = `checkbox-list-secondary-label-${img.id}`;
        console.log(img)
        return (
          <ListItem
            key={img.id}
            style={{justifyContent: "center", paddingRight: "10px", paddingLeft: "10px"}}
            onClick={() => handleClickOpen(`data:image/png;base64,${img.content}`)}
            disablePadding
          >
            <ListItemButton>
              <ListItemAvatar>
                <Avatar
                  src={`data:image/png;base64,${img.content}`}
                />
              </ListItemAvatar>
              <ListItemText style={{justifyContent: "center"}} id={labelId} primary={"SentryId: " + img.sentry_id} secondary={"ImageId: " + img.id}/>
              {img.timestamp}
            </ListItemButton>
          </ListItem>
        );
      })}
    </List>
    </React.Fragment>
  )
}