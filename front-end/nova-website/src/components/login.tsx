import { Modal, Box, Typography, TextField, Button } from "@mui/material";
import { BACKEND_SRC } from "../api/helper";
import axios from "axios";
import * as React from "react";
import UserState from "../userState";
import { Link, useNavigate } from "react-router-dom";

const style = {
    position: 'absolute' as 'absolute',
    top: '30%',
    left: '50%',
    transform: 'translate(-50%, -50%)',
    width: 400,
    bgcolor: 'background.paper',
    border: '2px solid #000',
    boxShadow: 24,
    p: 4,
    textAlign: 'center'
};

export default function Login() {

    return (
        <Modal
            open={true}
            aria-labelledby="modal-modal-title"
            aria-describedby="modal-modal-description"
        >
            <Box sx={style}>
                <Typography id="modal-modal-title" variant="h6" component="h2">
                    Login
                </Typography>
                <LoginForm />
                <Typography>or</Typography>
                <Link to="/sign-up">
                    <Button
                        variant="contained"
                        color="primary"
                    >
                        Sign Up
                    </Button>
                </Link>
            </Box>
        </Modal>
    );
}

function LoginForm() {

    const [email, setEmail] = React.useState('');
    const [pw, setPW] = React.useState('');
    let navigate = useNavigate();


    const handleEmailChange = (event) => {
        setEmail(event.target.value)
    };

    const handlePwChange = (event) => {
        setPW(event.target.value)
    };

    const loginSubmit = ((event) => {
        event.preventDefault();
        axios.post(`${BACKEND_SRC}auth/login`, {
            email: email,
            password: pw
        }).then((response) => {
            let state = UserState.getInstance().state;
            state["jwt_auth_token"] = response.data["token"];
            state["id"] = response.data["account_id"];
            navigate("/", {replace: true});
        }).catch((error) => {
            alert("Login failed");
        });
    });

    return (
        <Box
            component="form"
            sx={{
                '& .MuiTextField-root': { m: 1, width: '25ch' }
            }}
            noValidate
            autoComplete="off"
        >
            <TextField
                required
                id="email-field"
                label="Email"
                value={email}
                onChange={handleEmailChange}
            />
            <TextField
                id="outlined-password-input"
                label="Password"
                type="password"
                autoComplete="current-password"
                onChange={handlePwChange}
            />
            <br></br>
            <Button
                variant="contained"
                color="primary"
                onClick={loginSubmit}
                sx={{ marginTop: '1rem' }}
            >
                Enter
            </Button>
        </Box>
    );
}