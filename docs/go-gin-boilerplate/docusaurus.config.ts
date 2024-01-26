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
  baseUrl: "/ggb/",

  // GitHub pages deployment config.
  // If you aren't using GitHub pages, you don't need these.
  organizationName: "sanjayheaven", // Usually your GitHub org/user name.
  projectName: "ggb", // Usually your repo name.

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
          // editUrl: "https://github.com/sanjayheaven/ggb/issure",
        },
        blog: {
          showReadingTime: true,
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          editUrl:
            "https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/",
        },
        theme: {
          customCss: ["./src/css/custom.css", "./src/css/tailwind.css"],
        },
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
      { name: "title", content: "Go Gin Boilerplate" },
      {
        name: "description",
        content: "A development boilerplate based on the Gin framework",
      },
      { name: "keywords", content: "go gin boilerplate go-gin-boilerplate" },

      { property: "og:title", content: "Go Gin Boilerplate" },
      { property: "og:image", content: "img/cover.png" },
      // { property: "og:type", content: "Go Gin Boilerplate" },

      { name: "twitter:card", content: "summary" },

      // { name: "og:url", content: "https://github.com/sanjayheaven" },
    ],
    image: "img/cover.png",
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
              href: "https://github.com/sanjayheaven/ggb",
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
              href: "https://github.com/sanjayheaven/ggb/issues/new",
              label: "🆙 Help Improve",
              target: "_blank",
            },
          ],
        },
        {
          href: "https://github.com/sanjayheaven/ggb",
          position: "right",
          className: "header-github-link",
          "aria-label": "GitHub repository",
        },
      ],
    },
    footer: {
      style: "dark",
      links: [
        {
          title: "Docs",
          items: [
            { label: "Tutorial", to: "/docs/intro" },
            {
              label: "Swagger",
              to: "https://ggb.gganbu.services/swagger/index.html",
            },
          ],
        },
        {
          title: "Community",
          items: [
            {
              label: "GitHub",
              href: "https://github.com/sanjayheaven/ggb",
            },
            // {
            //   label: "Discord",
            //   href: "https://discordapp.com/invite/docusaurus",
            // },
          ],
        },
        {
          title: "More",
          items: [
            {
              label: "Dorvan's Blog",
              to: "https://blog.hdxsanjay.com",
            },
            // {
            //   label: "GitHub",
            //   href: "https://github.com/facebook/docusaurus",
            // },
          ],
        },
      ],
      copyright: `Copyright © ${new Date().getFullYear()} Dorvan, Inc. Built with Docusaurus.`,
    },
    prism: {
      theme: prismThemes.github,
      darkTheme: prismThemes.dracula,
    },
  } satisfies Preset.ThemeConfig,
};

export default config;
