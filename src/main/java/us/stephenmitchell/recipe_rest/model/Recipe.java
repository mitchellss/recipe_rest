package us.stephenmitchell.recipe_rest.model;

import com.fasterxml.jackson.annotation.JsonFormat;

import java.util.Date;

import javax.persistence.*;

@Entity
@Table(name = "recipe")
public class Recipe {

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    @Column(name="id")
    private Long id;

    @Column(name="title")
    private String title;

    @Column(name="recipe_datetime")
    //@JsonFormat(pattern="yyyy-MM-dd HH:mm:ss")
    private String datetime;

    public Recipe() {
        super();
    }

    public Recipe(String title, String datetime) {
        super();
        this.title = title;
        this.datetime = datetime;
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public String getDatetime() {
        return datetime;
    }

    public void setDatetime(String datetime) {
        this.datetime = datetime;
    }

    @Override
    public String toString() {
        return "RecipeModel{" +
                "id=" + id +
                ", title='" + title + '\'' +
                ", datetime='" + datetime + '\'' +
                '}';
    }
}
