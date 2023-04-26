import React, {useState} from 'react';
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import './App.css';
import HomePage from "./pages/HomePage";
import LandingPage from "./pages/LandingPage";

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
                  </Routes>
                )}
            </Router>
        </div>
    );
}

export default App;
