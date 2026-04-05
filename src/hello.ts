#!/usr/bin/env node
import chalk from 'chalk'
import { program } from 'commander'

program
  .name('hello')
  .argument('<name>', 'your name')
  .action((name: string) => {
    console.log(chalk.green(`Hello, ${name}!`))
  })

program.parse()