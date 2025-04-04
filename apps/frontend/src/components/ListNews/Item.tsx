import { Link } from 'react-router-dom';
import styles from './Styles.module.css';
import { INews } from '../../interface/INews';

const Item: React.FC<INews> = (props) => {
  return (
    <li key={props.id}>
      <div className={styles.card}>
        <Link to={`/${props.voteUrl}`}>
          <div className={styles.votearrow} title="upvote"></div>
        </Link>
        <div className={styles.content}>
          <div className={styles.title}>
            <Link to={`/${props.titleUrl}`}>
              <span className={styles.titleText}>{props.title}</span>
            </Link>
            <Link to={`/${props.siteUrl}`} className={`${!props.siteUrl ? styles.hidden : ''}`}>
              <span className={styles.siteText}>({props.site})</span>
            </Link>
          </div>
          <div className={styles.subTitle}>
            <span className={styles.siteText}>{props.score}</span>
            <span className={styles.siteText}>by</span>
            <Link to={`/${props.byUrl}`} className={!props.byUrl ? styles.hidden : ''}>
              <span className={styles.siteText}>{props.byUser}</span>
            </Link>
            <Link to={`/${props.ageUrl}`} className={!props.ageUrl ? styles.hidden : ''}>
              <span className={styles.siteText}>{props.age}</span>
            </Link>
            <span className={`${styles.siteText} ${!props.hideUrl ? styles.hidden : ''}`}>|</span>
            <Link to={`/${props.hideUrl}`} className={!props.hideUrl ? styles.hidden : ''}>
              <span className={styles.siteText}>hide</span>
            </Link>
            <span className={`${styles.siteText} ${!props.pastUrl ? styles.hidden : ''}`}>|</span>
            <Link to={`/${props.pastUrl}`} className={!props.pastUrl ? styles.hidden : ''}>
              <span className={styles.siteText}>past</span>
            </Link>
            <span className={`${styles.siteText} ${!props.comments ? styles.hidden : ''}`}>|</span>
            <Link to={`/${props.commentsUrl}`} className={!props.comments ? styles.hidden : ''}>
              <span className={styles.siteText}>{props.comments}</span>
            </Link>
          </div>
        </div>
      </div>
    </li>
  );
};

export default Item;
