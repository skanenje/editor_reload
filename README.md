# texEdit_tool

## _The Best Text Correction tool, Ever_
![go](https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Aqua.png)
 
 texEdit_tool is a text manipulation tool, offline-storage compatible, go powered text editor tool 
 
It is a tool with a variety of capabilities, it deals with some cases where the the text contains certain instances of tokens that are used to signify commands  to be undertaken on the preceding word or the text word.

Some of the tokens that are covered in the scope of this program include

+ Every instance of `(hex)` replaces the word before it to is actual value, in this case the word has been expressed in its hexadecimal form
+ Every instance of `(bin)` replaces the word before with the decimal version of the word 
+ Every instance of `(up)` converts the word before with the Uppercase version of it. if a number appears next to it, it turns the previously specified number of words accordingly.

+ Every instance of `(cap)` converts the word before with the Uppercase version of it. if a number appears next to it, it turns the previously specified number of words accordingly
+ Every instance of `(low)` converts the word before with the Uppercase version of it. if a number appears next to it, it turns the previously specified number of words accordingly
+ Every instance of the punctuations `.`, `,`, `!`, `?`, `:` and `;` is brought close to the previous word and with space apart from the next one.
+ Every instance of `a` is  turned into `an` if the next word begins with a vowel
+ The punctuation mark `'` will always be found with another instance of it and they are placed to the right and left of the word in the middle of them, without any spaces
## instalation and usage
texEdit requires [go](https://go.dev/dl/)  v 1.22.2 to run

Once installed  clone into this repository to do this 

- Open terminal and run
``` sh
git clone https://github.com/011369/starter
cd go-reloaded
ls go-reloaded
```
what you do now is copy the contents you want to edit into the sample file 
+ To do this, go to the location of the go-reloaded folder 
``` 
then hover above the icon then right click the drop down menu and select "open 
in terminal
once opened in vscode open the sample.txt file and copy the contents of the text you want to edit to it.
save and simply run;
go run . sample.txt result.txt
```

Now everything is setup you can use the tool anytime to do correction to your text.

Thank you and Happy coding

## HOW TO CONTRIBUTE ? 👷 

**1.** Fork [this](https://github.com/011369/starter) repository.

**2.** Clone the forked repository.

```terminal
git clone https://github.com/011369/starter
```

**3.** Navigate to the project directory.

```terminal
cd go-reloaded
```

**4.**  MAKE A NEW FOLDER WITH YOUR PROJECT NAME INSIDE 
<br>

**5.**  Also Add a README file in your project folder which consists of Description/screenshots about your project !
          
 
<br>

**6.** Create a new branch.

```terminal
git checkout -b <your_branch_name>
```

**7.** Add & Commit your changes.

```terminal
  git add .
  git commit -m "<your_commit_message>"
```

**7.** Push your local branch to the remote repository.

```terminal
git push -u origin <your_branch_name>
```

**8.** Create a Pull Request!

**Congratulations!** Sit and relax till we review your PR, you've made your contribution to (https://github.com/011369/starter) project
