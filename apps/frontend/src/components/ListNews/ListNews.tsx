import { INews } from '../../interface/INews';
import styles from './Styles.module.css';
import Item from './Item';
import ItemComment from './ItemComment';
import { IComment } from '../../interface/IComment';

interface Props {
  type?: 'comment' | 'news' | undefined;
  data: INews[] | IComment[];
  showMore?: boolean;
  onMoreClick?: () => void;
}

const ListNews: React.FC<Props> = (props) => {
  return (
    <div style={{ display: 'flex', flexDirection: 'column' }}>
      <ol>
        {props.data.map((list, idx) =>
          props?.type === 'comment' ? (
            <ItemComment {...(list as IComment)} key={idx} />
          ) : (
            <Item {...(list as INews)} key={idx} />
          )
        )}
      </ol>
      <button
        onClick={props?.onMoreClick}
        className={styles.more}
        style={{ display: !props?.showMore ? 'none' : 'block' }}>
        More
      </button>
    </div>
  );
};

export default ListNews;
