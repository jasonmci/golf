import { Player } from './models/player';
import { Course } from './models/course';
import { rollDice } from './utils/dice';
import readlineSync = require('readline-sync');
import * as fs from 'fs';
import * as path from 'path';

const loadCourseData = (filePath: string): Course => {
    const fullPath = path.join(__dirname, filePath);
    const data = fs.readFileSync(fullPath, 'utf-8');
    const jsonData = JSON.parse(data);

    return new Course(jsonData.name, jsonData.holes);
}

const main = () => {
    console.log("Welcome to TS Golf");

    const playerName = readlineSync.question('Enter player name: ');
    const player = new Player(playerName);

    const course = loadCourseData('../courses/pebble-beach.json');
    // get the player data next.
    console.log(`Player: ${playerName}, is playing ${course.name}`);

    // what hole number are you playing?
    const holeNumber = readlineSync.question('Enter hole number to view outcomes: ');
    course.printHoleOutcomes(holeNumber);

    const diceRoll = rollDice();
    console.log(`You rolled a ${diceRoll}`);

}

main();
