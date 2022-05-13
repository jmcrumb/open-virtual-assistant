import { AccountCircle } from '@mui/icons-material';
import { IconButton } from '@mui/material';
import * as React from 'react';
import { Link } from 'react-router-dom';


function Navbar() {
	return (
		<div className="Navbar">
			<Link to="/">
				<h1>NOVA</h1>
			</Link>
			<div className="account">
				<IconButton
				size="large"
				aria-label="account of current user"
				aria-controls="primary-search-account-menu"
				aria-haspopup="true"
				color="inherit"
				>
					<AccountCircle />
				</IconButton>
			</div>
		</div>
	);
}

export default Navbar;