import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import TimeAgo from 'javascript-time-ago';
import en from 'javascript-time-ago/locale/en.json';
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { Login } from "@mui/icons-material";
import { PluginViewPublic } from "./components/plugin";
import SignUp from "./components/signup";

TimeAgo.addDefaultLocale(en);

const root = ReactDOM.createRoot(
  document.getElementById("root")
);

root.render(
  <React.StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<App />}>
          <Route path="account" element={<App />}>
            <Route path="login" element={<Login />} />
            <Route path="sign-up" element={<SignUp />} />
          </Route>
          <Route path="plugin" element={<App />}>
            <Route path=":id" element={<PluginViewPublic />} />
          </Route>
          <Route path="*" element={
            <div>
              <p>There's nothing here!</p>
            </div>
          } />
        </Route>
      </Routes>
    </BrowserRouter>
  </React.StrictMode>
);