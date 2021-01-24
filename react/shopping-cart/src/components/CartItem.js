import React from "react";
import "./styles/CartItem.css"

export default function CartItem({
  id,
  name,
  price,
  quantity,
  updateQuantity,
}) {
  const subtract = () => updateQuantity(id, quantity - 1);
  const add = () => updateQuantity(id, quantity + 1);

  return (
    <div className="CartItem">
      <div>{name}</div>
      <div>{price}</div>
      <div>
        <button onClick={subtract} disabled={quantity === 0}>
          -
        </button>
        {quantity}
        <button onClick={add} disabled={quantity === 10}>+</button>
      </div>
      <div>Total: ${quantity * price}</div>
    </div>
  );
}
