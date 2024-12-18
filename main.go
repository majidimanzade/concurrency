package main

import faninfanout "github.com/majidimanzade/concurrency/fanin-fanout"

func main() {
	websitesFree := []string{
		"https://www.coursera.org/",
		"https://www.edx.org/",
		"https://www.khanacademy.org/",
		"https://www.udemy.com/",
		"https://www.codecademy.com/",
		"https://www.freecodecamp.org/",
		"https://www.youtube.com/",
		"https://www.udacity.com/",
		"https://www.futurelearn.com/",
		"https://www.skillshare.com/",
		"https://www.linkedin.com/learning/",
		"https://www.pluralsight.com/",
		"https://www.thegreatcoursesplus.com/",
		"https://www.openlearn.open.ac.uk/",
		"https://www.coursera.org/learn/learn-to-learn",
		"https://www.mit.edu/",
		"https://www.theoryofcomputing.org/",
		"https://www.ox.ac.uk/online-courses",
		"https://www.harvard.edu/",
	}

	pooliWebsites := []string{
		"https://www.masterclass.com/",
		"https://www.linkedin.com/learning/",
		"https://www.pluralsight.com/",
		"https://www.udacity.com/",
		"https://www.skillshare.com/",
		"https://www.teachable.com/",
		"https://www.codecademy.com/pro",
		"https://www.udemy.com/business/",
		"https://www.coursera.org/professional-certificates",
		"https://www.lynda.com/",
		"https://www.edx.org/professional-certificate",
		"https://www.datacamp.com/",
		"https://www.springboard.com/",
		"https://www.turing.com/",
		"https://www.groklearning.com/",
		"https://www.simplilearn.com/",
		"https://www.acadium.com/",
		"https://www.creativebug.com/",
		"https://www.skillwise.com/",
	}

	faninfanout.Crawl(websitesFree, pooliWebsites)
}
