import React from "react";
import "./styles.scss";
import Navbar from "./components/nav";
import Sandbox from "./components/sandbox";

const App: React.FC = () => {

  return (
      <div className="base">
        <Navbar />
        <Sandbox />
      </div>
  );
};

export default App;
