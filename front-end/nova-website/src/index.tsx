import React from "react";
import ReactDOM from "react-dom";
import App from "./App";
import TimeAgo from 'javascript-time-ago';
import en from 'javascript-time-ago/locale/en.json';
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Login from "./components/login";
import SignUp from "./components/signup";
import { PluginViewPublic } from "./components/plugin";

TimeAgo.addDefaultLocale(en)

const routes = (
  <React.StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<App />}>
          <Route path="account">
            <Route path="login" element={<Login />} />
            <Route path="sign-up" element={<SignUp />} />
          </Route>
          <Route path="plugin">
            <Route path=":id" element={<PluginViewPublic />}/>
          </Route>
        </Route>
      </Routes>
    </BrowserRouter>
    <App />
  </React.StrictMode>
);

ReactDOM.render(routes, document.getElementById("root"));
