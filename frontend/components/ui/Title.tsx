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
    <div>
      <h1 className={style}>{children}</h1>
    </div>
  )
}