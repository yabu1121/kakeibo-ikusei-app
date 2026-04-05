import { ReactNode } from "react";

type TitleProps = {
  type?: 'normal' | 'small' | 'large';
  children: ReactNode;
}

const TITLE_CONFIG = {
  normal: "font-bold text-2xl",
  small: "",
  large: "",
}

export const Title = ({ type = 'normal', children }: TitleProps) => {
  const style = TITLE_CONFIG[type]
  return (
    <div className="border-b pb-4">
      <h1 className={`${style} text-gray-800`}>{children}</h1>
    </div>
  )
}