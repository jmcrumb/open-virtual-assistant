import { Modal, Box, Typography, TextField, Button } from "@mui/material";
import { BACKEND_SRC } from "../api/helper";
import axios from "axios";
import * as React from "react";
import UserState from "../userState";

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

export default function SignUp() {

    return (
        <Modal
            open={true}
            aria-labelledby="modal-modal-title"
            aria-describedby="modal-modal-description"
        >
            <Box sx={style}>
                <Typography id="modal-modal-title" variant="h6" component="h2">
                    Sign Up
                </Typography>
                <SignUpForm />
            </Box>
        </Modal>
    );
}

function SignUpForm() {

    const [email, setEmail] = React.useState('');
    const [pw, setPW] = React.useState('');
    const [fName, setFName] = React.useState('');
    const [lName, setLName] = React.useState('');
    const [pwError, setPWError] = React.useState(false);


    const handleEmailChange = (event) => {
        setEmail(event.target.value)
    };

    const handlePwChange = (event) => {
        setPW(event.target.value)
    };

    const handleFNChange = (event) => {
        setFName(event.target.value)
    };

    const handleLNChange = (event) => {
        setLName(event.target.value)
    };

    const comparePasswords = (event) => {
        if (pw != event.target.value) {
            setPWError(true);
        }
    }

    const signUpSubmit = ((event) => {
        event.preventDefault();
        axios.post(`${BACKEND_SRC}account/`, {
            email: email,
            password: pw,
            first_name: fName,
            last_name: lName
        }).then((response) => {
            UserState.getInstance().state["id"] = response.data.id;
            axios.post(`${BACKEND_SRC}auth/login`, {
                email: email,
                password: pw
            }).then((response) => {
                UserState.getInstance().state["jwt_auth_token"] = response.data["token"];
            });
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
                id="outlined-required"
                label="Email"
                value={email}
                onChange={handleEmailChange}
            />
             <TextField
                required
                id="outlined-required"
                label="First Name"
                value={fName}
                onChange={handleFNChange}
            />
             <TextField
                required
                id="outlined-required"
                label="Last Name"
                value={lName}
                onChange={handleLNChange}
            />
            <TextField
                id="outlined-password-input"
                label="Password"
                type="password"
                autoComplete="current-password"
                onChange={handlePwChange}
            />
            <TextField
                id="outlined-password-input"
                label="Re-enter Password"
                type="password"
                autoComplete="current-password"
                error={pwError}
                onChange={comparePasswords}
            />
            <br></br>
            <Button
                variant="contained"
                color="primary"
                onClick={signUpSubmit}
                sx={{ marginTop: '1rem' }}
            >
                Enter
            </Button>
        </Box>
    );
}