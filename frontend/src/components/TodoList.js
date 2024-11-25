import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router";

export default function TodoList() {
  let navigate = useNavigate();
  const [todos, setTodos] = useState([]);
  useEffect(() => {
    fetch("http://localhost:5000/api/todos", {
      headers: {
        Accept: "application/json",
        "x-access-token": localStorage.getItem("token"),
      },
    })
      .then((res) => res.json())
      .then((json) => setTodos(json))
      .catch((err) => console.log(err));
  }, []);
  if (todos === undefined || todos === null || todos.length === 0) {
    return (
      <>
        <h3>No items</h3>
        <button style={styles.btnStyle} onClick={() => navigate("/create")}>
          Create
        </button>
      </>
    );
  }
  let todoItems = todos.map((todo) => (
    <tr style={styles.rowStyle}>
      <td key={todo.name} style={styles.colStyle}>
        {todo.name}
      </td>
      <td key={todo.description} style={styles.colStyle}>
        {todo.description}
      </td>
      <td key={todo.status} style={styles.colStyle}>
        {todo.status}
      </td>
      <td key={todo.createdBy} style={styles.colStyle}>
        {todo.createdBy}
      </td>
      <td key={todo.createdOn} style={styles.colStyle}>
        {todo.createdOn}
      </td>
    </tr>
  ));
  return (
    <div>
      <h3>Todos</h3>
      <table style={styles.tableStyle}>
        <thead>
          <tr style={styles.headStyle}>
            <th style={styles.colStyle}>Name</th>
            <th style={styles.colStyle}>Description</th>
            <th style={styles.colStyle}>Status</th>
            <th style={styles.colStyle}>Created By</th>
            <th style={styles.colStyle}>Created On</th>
          </tr>
        </thead>
        <tbody>{todoItems}</tbody>
      </table>
      <button style={styles.btnStyle} onClick={() => navigate("/create")}>
        Create
      </button>
    </div>
  );
}

const styles = {
  tableStyle: {
    "table-layout": "fixed",
    width: "100%",
    "border-collapse": "collapse",
    border: "1px solid purple",
    "text-align": "center",
  },
  rowStyle: {
    "border-bottom": "1px solid purple",
  },
  headStyle: {
    "border-bottom": "1px solid purple",
    backgroundColor: "#121010",
    color: "#ffffff",
  },
  colStyle: {
    "border-right": "1px solid purple",
    padding: "5px 0",
  },
  btnStyle: {
    marginTop: "15px",
  },
};
