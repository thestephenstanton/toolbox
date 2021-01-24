import CartItem from "./CartItem";
import "./styles/Cart.css";

export default function Cart({ items, updateItems }) {
  const updateQuantity = (id, newQuantity) => {
    const newItems = items.map((item) => {
      if (item.id === id) {
        return { ...item, quantity: newQuantity };
      }

      return item;
    });

    updateItems(newItems);
  };

  const grandTotal = items
    .reduce((total, item) => total + item.quantity * item.price, 0)
    .toFixed(2);

  return (
    <div className="Cart">
      <h1 className="Cart__title">Cart</h1>
      <div className="Cart__items">
        {items.map((item) => (
          <CartItem key={item.id} updateQuantity={updateQuantity} {...item} />
        ))}
      </div>
      <h2 className="Cart__total">Grand Total: ${grandTotal}</h2>
    </div>
  );
}
