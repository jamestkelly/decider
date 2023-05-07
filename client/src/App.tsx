import React, {useState} from 'react';
import {BrowserRouter as Router, Route, Routes} from "react-router-dom";
import './App.css';
import HomePage from "./pages/HomePage";
import LandingPage from "./pages/LandingPage";
import AboutPage from "./pages/AboutPage";
import RegisterPage from "./pages/RegisterPage";
import LoginPage from "./pages/LoginPage";
import SupportPage from "./pages/SupportPage";
import DocumentationPage from "./pages/DocumentationPage";


interface Props {
  setIsAuthenticated: (value: boolean) => void;
}

function App() {
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  return (
    <div className="App">
      <Router>
        {false ? (
          <Routes>
            <Route path={"/"} element={<LandingPage/>}/>
          </Routes>
        ) : (
          <Routes>
            <Route path={"/"} element={<HomePage/>}/>
            <Route path={"/about"} element={<AboutPage/>}/>
            <Route path={"/login"} element={<LoginPage/>}/>
            <Route path={"/register"} element={<RegisterPage/>}/>
            <Route path={"/support"} element={<SupportPage/>}/>
            <Route path={"/documentation"} element={<DocumentationPage/>}/>
          </Routes>
        )}
      </Router>
    </div>
  );
}

export default App;
