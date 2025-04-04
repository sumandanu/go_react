import styles from './Layout.module.css';
import { FC, ReactElement } from 'react';

interface LayoutInfoProps {
  children: ReactElement;
}

const LayoutInfo: FC<LayoutInfoProps> = ({ children }) => {
  return (
    <div className={styles.container}>
      <main>{children}</main>
    </div>
  );
};

export default LayoutInfo;
