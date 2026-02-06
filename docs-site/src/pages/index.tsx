import type { ReactNode } from 'react';
import clsx from 'clsx';
import Link from '@docusaurus/Link';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import Layout from '@theme/Layout';
import CodeBlock from '@theme/CodeBlock';
import HomepageFeatures from '@site/src/components/HomepageFeatures';
import Heading from '@theme/Heading';

import styles from './index.module.css';

const heroSnippet = `app:
  name: doit
  version: 0.1.0
  description: A task runner written in Go
  main_file: main.go
  authors:
    - Tidjee
  repo_url: https://github.com/tidjee-dev/doit

env:
  BIN_DIR: bin

tasks:
  deps:
    category: Dependencies
    description: Install dependencies
    commands:
      - go mod tidy

  build:
    category: Build
    description: Compile binary
    depends_on:
      - deps
    commands:
      - go build -o {{ .Env.BIN_DIR }}/doit {{ .App.MainFile }}`;

function HomepageHeader() {
  const { siteConfig } = useDocusaurusContext();
  return (
    <header className={clsx('hero hero--primary', styles.heroBanner)}>
      <div className={styles.heroBackdrop} />
      <div className="container">
        <div className={styles.heroGrid}>
          <div className={styles.heroCopy}>
            <p className={styles.kicker}>Go CLI Task Runner</p>
            <Heading as="h1" className={styles.heroTitle}>
              Ship repeatable workflows without magic.
            </Heading>
            <p className={styles.heroSubtitle}>{siteConfig.tagline}</p>
            <div className={styles.buttons}>
              <Link className="button button--lg button--primary" to="/docs/quick-start">
                Start in 5 minutes
              </Link>
              <Link className="button button--lg button--outline button--secondary" to="/docs/cli/reference">
                CLI Reference
              </Link>
            </div>
          </div>

          <div className={styles.heroCode}>
            <CodeBlock language="yaml" title="tasks.yml" className={styles.heroCodeBlock}>
              {heroSnippet}
            </CodeBlock>
          </div>
        </div>
      </div>
    </header>
  );
}

export default function Home(): ReactNode {
  const { siteConfig } = useDocusaurusContext();
  return (
    <Layout
      title={`${siteConfig.title} documentation`}
      description="Documentation for doit, the explicit task runner written in Go and automation workflows.">
      <HomepageHeader />
      <main>
        <HomepageFeatures />
      </main>
    </Layout>
  );
}
