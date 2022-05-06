import Sandbox from "./components/sandbox";
import React from "react";
import "./styles.scss";
import Navbar from "./components/nav";
import { BrowserRouter } from "react-router-dom";

const App: React.FC = () => {

  return (
    <div className="base">
      <Navbar />
    </div>
  );
};

export default App;
