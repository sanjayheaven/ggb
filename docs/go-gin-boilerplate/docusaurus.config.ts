import { themes as prismThemes } from "prism-react-renderer";
import type { Config } from "@docusaurus/types";
import type * as Preset from "@docusaurus/preset-classic";
import { EnumChangefreq } from "sitemap";

const config: Config = {
  title: "Go Gin Boilerplate",
  tagline:
    "一个基于 Gin 框架的开发脚手架，旨在帮助开发者快速搭建和开发 Web 应用程序。",
  favicon: "img/golang.png",

  // Set the production url of your site here
  url: "https://sanjayheaven.github.io/",
  // Set the /<baseUrl>/ pathname under which your site is served
  // For GitHub pages deployment, it is often '/<projectName>/'
  baseUrl: "/go-gin-boilerplate/",

  // GitHub pages deployment config.
  // If you aren't using GitHub pages, you don't need these.
  organizationName: "sanjayheaven", // Usually your GitHub org/user name.
  projectName: "go-gin-boilerplate", // Usually your repo name.

  onBrokenLinks: "ignore",
  onBrokenMarkdownLinks: "warn",

  // Even if you don't use internationalization, you can use this field to set
  // useful metadata like html lang. For example, if your site is Chinese, you
  // may want to replace "en" with "zh-Hans".
  i18n: {
    defaultLocale: "zh-Hans",
    locales: ["zh-Hans", "en"],
  },

  presets: [
    [
      "@docusaurus/preset-classic",
      {
        docs: {
          sidebarPath: "./sidebars.ts",
          // showLastUpdateTime: true,
          // showLastUpdateAuthor: true,
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          // editUrl: "https://github.com/sanjayheaven/go-gin-boilerplate/issure",
        },
        blog: {
          showReadingTime: true,
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          editUrl:
            "https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/",
        },
        theme: { customCss: "./src/css/custom.css" },
        sitemap: {
          changefreq: EnumChangefreq.WEEKLY,
          priority: 0.5,
          filename: "sitemap.xml",
        },
        gtag: { trackingID: "G-B19Q9X519X" },
      } satisfies Preset.Options,
    ],
  ],

  themeConfig: {
    // Replace with your project's social card
    metadata: [
      { name: "keywords", content: "gin boilerplate go-gin-boilerplate" },
      { name: "twitter:card", content: "summary" },
    ],
    image: "img/docusaurus-social-card.jpg",
    navbar: {
      title: "Go Gin Boilerplate",
      logo: { alt: "Logo", src: "img/golang.png" },
      items: [
        {
          type: "docSidebar",
          sidebarId: "tutorialSidebar",
          position: "left",
          label: "使用文档",
        },
        // { to: "/blog", label: "Blog", position: "left" },
        // { type: "localeDropdown", position: "right" },
        {
          label: "Support",
          position: "right",
          items: [
            {
              href: "https://github.com/sanjayheaven/go-gin-boilerplate",
              label: "🌟 Star on GitHub",
              target: "_blank",
            },
            {
              href: "https://ko-fi.com/J3J1T95FG",
              label: "☕️ Buy Me a Ko-fi",
              target: "_blank",
            },
            {
              href: "https://www.buymeacoffee.com/dorvan",
              label: "☕️ Buy Me a Coffee",
              target: "_blank",
            },
            {
              href: "https://github.com/sanjayheaven/go-gin-boilerplate/issues/new",
              label: "🆙 Help Improve",
              target: "_blank",
            },
          ],
        },
        {
          href: "https://github.com/sanjayheaven/go-gin-boilerplate",
          position: "right",
          className: "header-github-link",
          "aria-label": "GitHub repository",
        },
      ],
    },
    // footer: {
    //   style: "dark",
    //   links: [
    //     {
    //       title: "Docs",
    //       items: [
    //         {
    //           label: "Tutorial",
    //           to: "/docs/intro",
    //         },
    //       ],
    //     },
    //     {
    //       title: "Community",
    //       items: [
    //         {
    //           label: "Stack Overflow",
    //           href: "https://stackoverflow.com/questions/tagged/docusaurus",
    //         },
    //         {
    //           label: "Discord",
    //           href: "https://discordapp.com/invite/docusaurus",
    //         },
    //         {
    //           label: "Twitter",
    //           href: "https://twitter.com/docusaurus",
    //         },
    //       ],
    //     },
    //     {
    //       title: "More",
    //       items: [
    //         {
    //           label: "Blog",
    //           to: "/blog",
    //         },
    //         {
    //           label: "GitHub",
    //           href: "https://github.com/facebook/docusaurus",
    //         },
    //       ],
    //     },
    //   ],
    //   copyright: `Copyright © ${new Date().getFullYear()} My Project, Inc. Built with Docusaurus.`,
    // },
    prism: {
      theme: prismThemes.github,
      darkTheme: prismThemes.dracula,
    },
  } satisfies Preset.ThemeConfig,
};

export default config;
