## A Practical Guide To Approach For Creating URL Shortner Application ##

### What is a Short URL? ###

<p>
A Short URL refers to a web service that transforms a long URL into a much shorter version.This helps to simplify the complex URLs that may be difficult to share or fit into a character limited platforms like sms/tweets.In addition, it is easy and less error prone to type a short url when compared to its longer version.  
</p>

### What are the benefits of a Short URL? ###

<p> 

<ol>

<li> Easy to share and remember. </li>
<li> Helpful to track the clicks on URLs for Analytics purposes. </li>
<li> Helps to Beautify any Long URLs. </li>

</ol>

</p>

### Requirements to create a Short URL Application? ###

<p>

<ul>

<li>Generate a short url with the given original url.</li>
<li>The short url (or link) should redirect users to original url (link).</li>
<li>Provide option to create custom short url as given by end user.</li>

</ul>

</p>

### How a Short URL looks like? ###

<p>
Commomly the most URL shortening application or services provide the short URL where:<br/>
Example: `http://localhost:9999/1L9zO9P`

- Fist Part: is the domain name of the service. example: `localhost:9999`
- Second Part: is a string formed by a set of random characters. example `1L9zO9P`<br/>
This random string should be unique for each short URL created for a long URL.

- Length of Short URL or Random String:<br/>

The unique random characters can be generated using either `base62` or `md5`. <br/>
Both `base62` and `MD5` algorithm outputs consist of 62 character combinations `(a-z A-Z 0-9)`.<br/>

`Base62` takes integer type as input. So we need to first convert the long URL to a random number then pass this random number to the base62 algorithm. <br/>

`MD5` takes string as the input and generates a random fixed length string output, we can directly pass the long URL as input.<br/>

</p>

### URL Shortnening Algorithms? ###

<ul>

<li>
Hashing Algorithm: Use a hashing algorithm such as MD5, SHA-256, or CRC32 to generate a hash value from the long URL. The hash value can then be encoded using base62 or base64 encoding to create a short URL.
</li>
<br/>

<li>
Base Conversion Algorithm: Convert the unique ID of the long URL into a shorter representation using base conversion techniques. For example, you can convert the decimal representation of the ID into a base58 or base62 encoding, excluding easily confused characters like ‘0’, ‘O’, ‘1’, ‘I’, etc.
</li>
<br/>

<li>
Bijective Function Algorithm: Utilize a bijective function that maps a unique identifier to a short URL and vice versa. For example, you can use a function like the Bijective Conversion Function (BCF) algorithm, which converts a decimal ID into a sequence of characters using a predefined set of characters.
</li>
<br/>

<li>
Randomized Approach: Generate a random sequence of characters of a fixed length (e.g., 6 or 8 characters) to create the short URL. Although this approach may not guarantee uniqueness, you can perform a lookup in the database to ensure uniqueness and regenerate if there’s a collision.
</li>
<br/>

<li>
Counter Based Approach: When the server gets a request to convert a long URL to short URL, it first talks to the counter to get a count value.This value is then passed to the `base62` algorithm to generate random string. Making use of a counter to get the integer value ensures that the number we get will always be unique by default because after every request the counter increments its value.
</li>

</ul>

### How `base62` Algorithm Works? ###

```text

The Base62 encoding scheme is a binary-to-text encoding scheme that represents binary data in an ASCII string format. It uses character that can be one of the following:

A lower case alphabet [‘a’ to ‘z’] -> total 26 characters
An upper case alphabet [‘A’ to ‘Z’] -> total 26 characters
A digit [‘0′ to ‘9’] -> total 10 characters

Hence, there are total 26 + 26 + 10 = 62 possible characters. [base62].

The typical order is 0-9 for numbers (values 0 to 9), A-Z for uppercase letters (values 10-35), a-z for lowercase letters (values 36-61)

In this application, it is decided to go with random character string of length 7, so we can have 62^7 ~ 3.5 trillion combinations to work with.This is more than enough for a sample application.

```

### Conversion to `base62`? ###

<ul>

<li>Divide the number by 62.</li>
<li>Record the remainder as the rightmost digit.</li>
<li>Use the quotient as the new number to be divided.</li>
<li>Repeat the process until the number is zero.</li>
<li>The sequence of remainders (read from bottom to top) forms the Base62 encoded string.</li>

</ul>


