import React from "react";
import "./styles.scss";
import "./components/Publish.css";
import "./components/PluginList.css";
import "./components/PluginPreview.css";
import "./components/Rating.scss";
import "./components/Home.css";
import "./components/Navbar.css";
import Navbar from "./components/Navbar";
// import Sandbox from "./components/sandbox";
// import Login from "./components/Login";
// import { PluginViewPublic } from "./components/PluginViewPublic";
// import SignUp from "./components/signup";
import { Routes, Route, useNavigate } from "react-router-dom";
import Home from "./components/Home";
// import PluginSearch from "./components/PluginSearch";
import PublishPlugin from "./components/Publish";
// import ProfileView from "./components/ProfileView";
import { GlobalStateContext } from "./globalState";
import { Login } from "@mui/icons-material";
import PublishedPlugins from "./components/PublishedPlugins";



const App: React.FC = () => {

  const navigate = useNavigate();
  const context = React.useContext(GlobalStateContext);

  React.useEffect(() => {
      if (context.id == "") {
        navigate("/login", { replace: true });
      }
    }, []);

  return (
      <div className="base">
		<Navbar />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<Login />} />
          {/* <Route path="/sign-up" element={<SignUp />} />
          <Route path="/profile/:id" element={<ProfileView />} />
          <Route path="/plugin/:id" element={<PluginViewPublic />} /> */}
          {/* <Route path="/plugin/search/:query" element={<PluginSearch />}></Route> */}
          <Route path="/plugin/published" element={<PublishedPlugins />}></Route>
          <Route path="/plugin/publish" element={<PublishPlugin />}></Route>
          <Route path="*" element={<h1>E404: Page Not found</h1>} />
        </Routes>
      </div>
  );
};

export default App;
