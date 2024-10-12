import React from "react";
import { useNavigate } from "react-router-dom";

const LandingPage = () => {
  const navigate = useNavigate();

  const goToSignup = () => {
    navigate("/signup");
  };

  const goToSignin = () => {
    navigate("/signin");
  };

  return (
    <div>
      <h1>Welcome to the App</h1>
      <button onClick={goToSignup}>Sign Up</button>
      <button onClick={goToSignin}>Sign In</button>
    </div>
  );
};

export default LandingPage;
