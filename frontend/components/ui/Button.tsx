import { ComponentPropsWithoutRef, ReactNode } from "react";

type ButtonProps = {
  variant?: "primary" | "secondary" | "danger";
  children: ReactNode;
} & ComponentPropsWithoutRef<"button">;

export const Button = ({ 
  type = "button", 
  variant = "primary", 
  className, 
  children, 
  ...props 
}: ButtonProps) => {
  
  const baseStyle = "px-4 py-2 rounded font-medium transition-colors focus:outline-none focus:ring-2";
  
  const variants = {
    primary: "bg-blue-600 text-white hover:bg-blue-700 focus:ring-blue-500",
    secondary: "bg-gray-200 text-gray-800 hover:bg-gray-300 focus:ring-gray-400",
    danger: "bg-red-600 text-white hover:bg-red-700 focus:ring-red-500",
  };

  return (
    <button 
      type={type} 
      className={`${baseStyle} ${variants[variant]} ${className}`} 
      {...props} 
    >
      {children}
    </button>
  );
};