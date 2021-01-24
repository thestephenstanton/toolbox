import { useEffect, useState } from "react";
import Cart from "./components/Cart";

const ITEMS = [
  { id: 1, name: "Taco Seasoning", price: 2.25, quantity: 2 },
  { id: 2, name: "Pinto Beans", price: 1.99, quantity: 3 },
  { id: 3, name: "Sour Cream", price: 3.5, quantity: 1 },
];

function App() {
  const cachedState = JSON.parse(localStorage.getItem("items"));
  const [items, setItems] = useState(cachedState || ITEMS);

  const updateItems = (updatedItems) => {
    setItems(updatedItems);
  };

  useEffect(() => {
    localStorage.setItem("items", JSON.stringify(items))
  }, [items])

  return (
    <div className="App">
      <Cart items={items} updateItems={updateItems} />
    </div>
  );
}

export default App;
