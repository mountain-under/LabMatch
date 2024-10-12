import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Signup from "./pages/Signup";
import Signin from "./pages/Signin";
import Home from "./pages/Home";
import LandingPage from "./pages/LandingPage";

const App = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<LandingPage />} />  {/* 最初に表示されるページ */}
        <Route path="/signup" element={<Signup />} />  {/* サインアップページ */}
        <Route path="/signin" element={<Signin />} />  {/* サインインページ */}
        <Route path="/home" element={<Home />} />      {/* サインイン後のホームページ */}
      </Routes>
    </Router>
  );
};

export default App;
