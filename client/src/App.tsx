import React, {useState} from 'react';
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
            {false ? (
                <LandingPage/>
            ) : (
                <HomePage/>
            )}
        </div>
    );
}

export default App;
