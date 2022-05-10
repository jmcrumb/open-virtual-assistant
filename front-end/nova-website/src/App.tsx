import React from "react";
import "./styles.scss";
import "./components/PublishPlugin.css";
import "./components/PluginList.css";
import "./components/PluginPreview.css";
import "./components/Rating.scss";
import "./components/Home.css";
import Navbar from "./components/Navbar";
import Sandbox from "./components/sandbox";
import Login from "./components/Login";
import { PluginViewPublic } from "./components/PluginViewPublic";
import SignUp from "./components/signup";
import { Routes, Route } from "react-router-dom";

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
