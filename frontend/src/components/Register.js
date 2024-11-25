import React, { useState } from "react";

export default function Register() {
  const [newuser, setNewUser] = useState();
  const handleRegister = (property, event) => {
    const userCopy = { ...newuser };
    userCopy[property] = event.target.value;
    setNewUser(userCopy);
  };
  const handleSubmit = (event) => {
    event.preventDefault();
    console.log(newuser.email);
    console.log(newuser.password);
    fetch("http://localhost:4000/api/users", {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        firstName: newuser.firstName,
        lastName: newuser.lastName,
        email: newuser.email,
        password: newuser.password,
      }),
    })
      .then((res) => {
        console.log(res);
      })
      .catch((err) => {
        console.log(err);
      });
  };
  return (
    <form>
      <fieldset>
        <legend>Register</legend>
        <p>
          <label htmlFor="firstName">
            First Name:
            <input
              type="text"
              name="firstName"
              required
              onChange={(event) => handleRegister("firstName", event)}
            />
          </label>
        </p>
        <p>
          <label htmlFor="lastName">
            Last Name:
            <input
              type="text"
              name="lastName"
              required
              onChange={(event) => handleRegister("lastName", event)}
            />
          </label>
        </p>
        <p>
          <label htmlFor="email">
            Email:
            <input
              type="text"
              name="email"
              required
              onChange={(event) => handleRegister("email", event)}
            />
          </label>
        </p>
        <p>
          <label htmlFor="password">
            Password:
            <input
              type="text"
              name="password"
              required
              onChange={(event) => handleRegister("password", event)}
            />
          </label>
        </p>
        <p>
          <input type="submit" value="Register" onClick={handleSubmit} />
        </p>
      </fieldset>
    </form>
  );
}
