import React from "react";
import "./styles.scss";
import Navbar from "./components/nav";
import Sandbox from "./components/sandbox";
import Login from "./components/login";
import { PluginViewPublic } from "./components/plugin";
import SignUp from "./components/signup";
import { BrowserRouter, Routes, Route } from "react-router-dom";

const App: React.FC = () => {

  return (
    <div className="base">
      <Navbar />
      <Routes>
        <Route path="/" element={<Sandbox />} />
        <Route path="/login" element={<Login />} />
        <Route path="/sign-up" element={<SignUp />} />
        <Route path="/plugin/:id" element={<PluginViewPublic />} />
        <Route path="*" element={<h1>E404: Page Not found</h1>} />
      </Routes>
    </div>
  );
};

export default App;
