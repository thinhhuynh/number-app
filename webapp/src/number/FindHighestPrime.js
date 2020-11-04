import React, { useState, useEffect } from 'react';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import Alert from '@material-ui/lab/Alert';

// import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import axios from 'axios';

const useStyles = makeStyles((theme) => ({
  image: {
    backgroundImage: 'url(https://source.unsplash.com/random)',
    backgroundRepeat: 'no-repeat',
    backgroundColor:
      theme.palette.type === 'light' ? theme.palette.grey[50] : theme.palette.grey[900],
    backgroundSize: 'cover',
    backgroundPosition: 'center',
  },
  paper: {
    margin: theme.spacing(8, 4),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  form: {
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}));

const FindHighestPrime = () => {
  const classes = useStyles();
  
  const [number, setNumber] = useState(null);
  const [highestPrime, setHighestPrime] = useState(null);
  const [message, setMessage] = useState('');
  const [type, setType] = useState('info');
  
  const handleChangeInput = e => {
    setNumber(e.target.value);
  }

  useEffect(() => {
    let type_ = 'info';
    let message_ = 'Input number then submit to find highest prime!';

    if (highestPrime === 'INVALID_INPUT_NUMBER') {
      type_ = 'error';
      message_ = 'Invalid input value. Please input number is a integer';
    }
    if (highestPrime === -1) {
      type_ = 'warning';
      message_ = `Not found highest prime less than ${number}`;
    }
    if (highestPrime > 1) {
      type_ = 'success';
      message_ = `Highest prime less than ${number} is ${highestPrime}`;
    }
    
    setType(type_);
    setMessage(message_);
  }, [highestPrime]);
  

  const handleFindHighestPrime = async () => {
    setHighestPrime(null);
    try {
      const resp = await axios.get(`${process.env.REACT_APP_API_URL}/api/find-highest-prime/?number=${number}`);
      setHighestPrime(resp.data);
    } catch(e) {
      setHighestPrime(e.response.data.error);
    }
  };

  return (
    <Grid container component="main" className={classes.root}>
      <Grid item xs={12} component={Paper} elevation={6} square>
        <div className={classes.paper}>
          <Typography component="h1" variant="h5">
            Find a highest prime number lower than input number
          </Typography>
          <form className={classes.form} noValidate>
            <TextField
              variant="outlined"
              margin="normal"
              required
              fullWidth
              id="number"
              label="Enter a number"
              name="number"
              autoComplete="number"
              autoFocus
              size="small"
              type="number"
              inputProps={
                {"min": "1"}
              }
              value={number}
              onChange={handleChangeInput}
            />
            <Button
              type="button"
              fullWidth
              variant="outlined"
              color="primary"
              className={classes.submit}
              size="large"
              disabled={!number}
              onClick={handleFindHighestPrime}
            >
              Find highest prime number
            </Button>

            <Grid container>
              <Grid item xs>
                <Alert severity={type}>{message}
                </Alert>
              </Grid>
            </Grid>            
          </form>
        </div>
      </Grid>
    </Grid>
  );
}

export default FindHighestPrime;