import { Avatar, Box, Button, Checkbox, CssBaseline, FormControlLabel, Grid2, Link, Paper, TextField, Typography } from '@mui/material';
import React, { useState } from 'react'
import Copyright from './Copyright';
import { LockOutlined } from '@mui/icons-material';

export default function Login() {
    const [account, setAccount] = useState({email:"",password:""});
    const handleLogin = (event) => {
        event.preventDefault();
        console.log(account.email);
        console.log(account.password);
        fetch('http://localhost:4000/api/login', {
            method: 'POST',
            headers: {
              'Accept': 'application/json',
              'Content-Type': 'application/json',
            },
            body: JSON.stringify({
              email: account.email,
              password: account.password,
            })
        }).then((res) => {
            console.log(res);
        }).catch((err) => {
            console.log(err);
        })
    }

    const handleAccount = (property, event)=>{
        const accountCopy = {...account};
        accountCopy[property] = event.target.value;
        setAccount(accountCopy);
    }
    return (
        <div>
            <Grid2 container component="main" className={styles.Container}>
                <CssBaseline />
                {/* <Grid item xs={false} sm={4} md={7} className={classes.image} /> */}
                <Grid2
                    item
                    xs={12}
                    sm={8}
                    md={5}
                    component={Paper}
                    elevation={1}
                    square
                >
                    <div>
                        <Avatar className={styles.avatar}>
                            <LockOutlined />
                        </Avatar>
                        <Typography component="h1" variant="h5">
                            Sign in
                        </Typography>
                        <form noValidate>
                            <TextField
                                onChange={(event)=>handleAccount("username", event)}
                                variant="outlined"
                                margin="normal"
                                required
                                fullWidth
                                id="username"
                                label="Username"
                                name="username"
                                autoFocus
                            />
                            <TextField
                                onChange={(event)=>handleAccount("password", event)}
                                variant="outlined"
                                margin="normal"
                                required
                                fullWidth
                                name="password"
                                label="Password"
                                type="password"
                                id="password"
                                autoComplete="current-password"
                            />
                            <FormControlLabel
                                control={<Checkbox value="remember" color="primary" />}
                                label="Remember me"
                            />
                            <Button
                                type="submit"
                                fullWidth
                                variant="contained"
                                color="primary"
                                onClick = {handleLogin}
                            >
                            Sign In
                            </Button>
                            <Grid2 container>
                                <Grid2 item>
                                    <Link href="#" variant="body2">
                                    {"Don't have an account? Sign Up"}
                                    </Link>
                                </Grid2>
                            </Grid2>
                            <Box mt={5}>
                                <Copyright />
                            </Box>
                        </form>
                    </div>
                </Grid2>
            </Grid2>
        </div>
    )
}

export const styles = {
    Container: {
        justifyContent: "center",
    },
    avatar: {
        justifyContent: "center",
    }
  }