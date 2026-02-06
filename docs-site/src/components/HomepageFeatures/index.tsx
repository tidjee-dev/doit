import type { ReactNode } from 'react';
import clsx from 'clsx';
import Heading from '@theme/Heading';
import styles from './styles.module.css';

type FeatureItem = {
  title: string;
  icon: string;
  description: ReactNode;
};

const FeatureList: FeatureItem[] = [
  {
    title: 'Deterministic Execution',
    icon: '01',
    description: (
      <>
        Tasks run in declared order with explicit dependency links. No hidden graph and no surprise parallel steps.
      </>
    ),
  },
  {
    title: 'Strict Config Contracts',
    icon: '02',
    description: (
      <>
        Keep `tasks.yml` readable and predictable with validated fields, naming rules, and one-dependency limits.
      </>
    ),
  },
  {
    title: 'Portable by Default',
    icon: '03',
    description: (
      <>
        A single Go binary, shell-based commands, and straightforward templates make adoption easy across CI and local dev.
      </>
    ),
  },
];

function Feature({ title, icon, description }: FeatureItem) {
  return (
    <div className={clsx('col col--4', styles.cardWrap)}>
      <article className={styles.card}>
        <div className={styles.icon}>{icon}</div>
        <Heading as="h3" className={styles.cardTitle}>
          {title}
        </Heading>
        <p className={styles.cardText}>{description}</p>
    </article>
      </div>
  );
}

export default function HomepageFeatures(): ReactNode {
  return (
    <section className={styles.features}>
      <div className="container">
        <Heading as="h2" className={styles.sectionTitle}>
          Why teams choose doit
        </Heading>
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
