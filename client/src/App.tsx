import React, {useState} from 'react';
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import './App.css';
import HomePage from "./pages/HomePage";
import LandingPage from "./pages/LandingPage";
import AboutPage from "./pages/AboutPage";

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
                  </Routes>
                )}
            </Router>
        </div>
    );
}

export default App;
